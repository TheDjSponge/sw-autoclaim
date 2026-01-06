# Infrastructure handling through ansible

To deploy on various devices (here raspberry pis) and ensure they're in the state we want, we use a program named ansible. That
way, we're able to easily ensure that the host entity is in an
expected and proper state. Ansible is used for the initial 
deployment of the docker-compose stack through the [compose_deploy](./roles/compose_deploy/) as well as launching [a service](./roles/service_provision/)
that periodically pulls new images on the registry and runs the 
stack again in case any new updated image is pushed on the registry.


## Install ansible and deploy on host device
```bash
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible -y
```

In order for ansible-playbooks to install and perform changes
on host devices, it will need a passwordless root-access 
through ssh (preferably) as well as some required information passed as variables. 
Read about [having an exploitable user on host](#create-passwordless-priviledged-user) 
as well as [proving variables](#providing-variables) before trying to run it.


To provide all the files and start the service, first cd into
this infra directory 
```bash
cd infra

ansible-playbook -i inventory/production provision.yml --vault-password-file ../vault.secret
```

## Ansible project structure

Ansible performs deployments using playbooks, yml scripts that
dictate what the resulting state of a target device should look
like. The main advantage is to avoid the burden of manually 
setting up dependencies, transferring files and so on.

Another main advantage of the idea behind ansible is to 
guarantee that the device is in a certain state through 
idempotent actions (performing the same actions twice
result in the same final state).

the structure is split between 3 types of "objects" in the 
ansible paradigm. 

- The **inventory** dictates what the host devices are,
their ip and what group they're in. For example here, we have
a production inventory that consists in a raspberry pi group
with a single device. For convenience purpose we also define
group specific vars under inventory items

- The **group_vars** dictates variables than can be used in playbooks or Jinja2 templates that are later rendered into
the host devices

- The **roles** define sequence of actions to execute on the 
host entities, each directory defines a set of actions that
guarantees a desired target state and may use some templates
or files to use in that process.

Finally, a playbook defines what roles should be played on 
what groups.

## Providing variables
Similarly to the .env file, we provide environment details through
the yml files found in group_vars directories. Sensitive 
information is provided through an encrypted vault and edited
with 
```bash
ansible-vault edit group_vars/vault.yml --vault-password-file ../vault.secret
```
where vault.secret is a file containing the vault(s) secret key (of course you'll need to create your own vaults if you don't have my secrets :D)

Here, we made the choice to define some sensitive variables as used by [all hosts](./group_vars/all/). This file contains the following vars: 
```yaml
# in group_vars/vault.yml
vault_registry_user: <github_package_registry_username>
vault_registry_pat: <github_package_registry_PAT>
```

and other variables specific to the production raspberry pi
```yaml
# in inventory/production/group_vars/raspberry_pi
vault_db_user: ""
vault_db_password: ""
vault_discord_bot_token: ""

vault_check_user_url: ""
vault_claim_coupon_url: ""
vault_gamecodes_url: ""

```

## Create a passwordless priviledged user
If you don't have one create an ssh key on the host user
```bash
ssh-keygen -t ed25519 -C "ansible-control" # If we don't have one
ssh [user]@[host_ip]
```

Deployment requires the target host to have a "deploy" user that 
contains sudo priviledge with ssh access

```bash
adduser deploy
usermod -aG sudo deploy
```

The user will also need ssh access and ideally paswordless sudo and make it unwritable
```bash
echo "deploy ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/deploy
chmod 440 /etc/sudoers.d/deploy # Give readonly perm for everyone
```

Then install ssh keys for the user
```bash
mkdir -p /home/deploy/.ssh
nano /home/deploy/.ssh/authorized_keys # Here add your machine's ssh pub key
chmod 700 /home/deploy/.ssh # give read/write/execute permissions to owner only 
chmod 600 /home/deploy/.ssh/authorized_keys # give read/write permissions to owner for autorized key
chown -R deploy:deploy /home/deploy/.ssh # set the deploy user and group as owner of the .ssh directory
```

