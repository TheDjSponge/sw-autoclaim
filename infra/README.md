To deploy on various devices (here raspberry pis) and ensure they're in the state we want, we use a program named ansible that is installed on the control machine with 

```bash
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible -y
```

Then we make sure that we have acces to all target devices with ssh
```bash
ssh-keygen -t ed25519 -C "ansible-control" # If we don't have one
ssh [user]@192.168.1.XX
```

Finally, we define targets with a hosts.ini file containing the following:

```ini
[pi_cluster]
# Replace with your Pi's actual IP
192.168.1.XX

[pi_cluster:vars]
# Replace 'pi' if you use a different username on the Raspberry Pi
ansible_user=[user]
# This tells Ansible to use the Python interpreter on the Pi
ansible_python_interpreter=/usr/bin/python3
```

check connections with 

```bash
ansible -i hosts.ini pi_cluster -m ping
```

Then, we can define some secrets that we'll provide to the target devices with

```bash
ansible-vault encrypt_string 'ghp_YOUR_TOKEN' --name 'ghcr_pat'
```

And finally define a playbook to ensure proper configuration

```yml
---
- name: Setup Pi for Deployment
  hosts: pi_cluster
  become: yes # Run as root (sudo)
  vars:
    # ---------------------------------------------------------
    # PASTE YOUR VAULT BLOCK HERE (Indent it correctly!)
    # ---------------------------------------------------------
    ghcr_pat: !vault |
          $ANSIBLE_VAULT;1.1;AES256
          663836343939336... (your encrypted block) ...366
    
    # Replace with your GitHub username
    gh_user: "your_github_username"

  tasks:
    - name: Ensure Docker is installed
      package:
        name: docker.io
        state: present

    - name: Log into GitHub Container Registry
      community.docker.docker_login:
        registry: ghcr.io
        username: "{{ gh_user }}"
        password: "{{ ghcr_pat }}"
        reauthorize: yes
```