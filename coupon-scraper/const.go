package main

const HTML string = `<div class="container">
  	<nav aria-label="breadcrumb">
      <ol class="breadcrumb box-shadow">
        <li class="breadcrumb-item"><a href="/">Dashboard</a></li>
        <li class="breadcrumb-item"><a>Tools</a></li>
        <li class="breadcrumb-item active" aria-current="page">Game Codes</li>
      </ol>
    </nav>
    
    <div class="card card-default box-shadow col-12 order-1 px-0">
		<div class="card-header pt-1 pb-1 pl-1 pr-1">
			<div class="row">
				<div class="col-9">
					Summoners War - Active Codes
				</div>
				<div class="col-3 text-right">
					<button type="button" class="btn btn-xs btn-primary" onClick="function_addGameCode();"><i class="fas fa-plus"></i></button>
					<button type="button" class="btn btn-xs btn-primary" onClick="function_viewModalFAQ(57);"><i class="fas fa-question"></i></button>
				</div>
			</div>
		</div>
   		<div class="card-body px-1 px-md-3">
   			
   			<div class="d-none d-md-block">
	   			<table id="activeGameCodes" class="table table-hover mb-0 table-sm font-size-08em w-100" cellspacing="0">
					<thead>
						<tr>
							<th class="text-center" style="width:20%;">Code</th>
							<th class="text-center" style="width:24%;">Date Added</th>
							<th class="text-center">Rewards</th>
							<th class="text-center" style="width:15%;">Like/Dislike</th>
						</tr>
					</thead>
					
					<tbody>
						
						<tr>
							<td>
								<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
								<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="SW2025DEC9PJ">SW2025DEC9PJ</a>
							</td>
							<td class="text-md-center">11/30/2025 03:08 PM UTC</td>
							<td><div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div></td>
							<td class="text-md-center " id="TD_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
							</td>
						</tr>
						
					</tbody>
				</table>
			</div>
			
				
			
			<div class="row d-md-none font-size-08em px-1">
				<div class="col-4 text-bold mb-1">Code:</div>
				<div class="col-8 mb-1">
					<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
					<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="SW2025DEC9PJ">SW2025DEC9PJ</a>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Date Added:</div>
				<div class="col-8 mb-1 mt-1">
					11/30/2025 03:08 PM UTC
				</div>
				
				<div class="col-4 text-bold mb-1 mt-2">Rewards:</div>
				<div class="col-8 mb-1 mt-1">
					<div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Like/Dislike:</div>
				<div class="col-8 mb-1 mt-1" id="DIV_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
				</div>
			</div>
			<hr class="d-md-none my-4"/>
				
   			
   		
	   		<div class="row mb-2">
				<div class="col-12 mt-2 px-1 text-center text-footnote">
					Data provided by the SWGT Community. SWGT not responsible for inaccuracies. Also available on SWGT homepage.
				</div>
			</div>
   		</div>
   	</div>

	<div class="container">
  	<nav aria-label="breadcrumb">
      <ol class="breadcrumb box-shadow">
        <li class="breadcrumb-item"><a href="/">Dashboard</a></li>
        <li class="breadcrumb-item"><a>Tools</a></li>
        <li class="breadcrumb-item active" aria-current="page">Game Codes</li>
      </ol>
    </nav>
    
    <div class="card card-default box-shadow col-12 order-1 px-0">
		<div class="card-header pt-1 pb-1 pl-1 pr-1">
			<div class="row">
				<div class="col-9">
					Summoners War - Active Codes
				</div>
				<div class="col-3 text-right">
					<button type="button" class="btn btn-xs btn-primary" onClick="function_addGameCode();"><i class="fas fa-plus"></i></button>
					<button type="button" class="btn btn-xs btn-primary" onClick="function_viewModalFAQ(57);"><i class="fas fa-question"></i></button>
				</div>
			</div>
		</div>
   		<div class="card-body px-1 px-md-3">
   			
   			<div class="d-none d-md-block">
	   			<table id="activeGameCodes" class="table table-hover mb-0 table-sm font-size-08em w-100" cellspacing="0">
					<thead>
						<tr>
							<th class="text-center" style="width:20%;">Code</th>
							<th class="text-center" style="width:24%;">Date Added</th>
							<th class="text-center">Rewards</th>
							<th class="text-center" style="width:15%;">Like/Dislike</th>
						</tr>
					</thead>
					
					<tbody>
						
						<tr>
							<td>
								<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
								<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="POUETPOUET">POUETPOUET</a>
							</td>
							<td class="text-md-center">11/30/2025 03:08 PM UTC</td>
							<td><div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div></td>
							<td class="text-md-center " id="TD_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
							</td>
						</tr>
						
					</tbody>
				</table>
			</div>
			
				
			
			<div class="row d-md-none font-size-08em px-1">
				<div class="col-4 text-bold mb-1">Code:</div>
				<div class="col-8 mb-1">
					<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
					<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="POUETPOUET">POUETPOUET</a>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Date Added:</div>
				<div class="col-8 mb-1 mt-1">
					11/30/2025 03:08 PM UTC
				</div>
				
				<div class="col-4 text-bold mb-1 mt-2">Rewards:</div>
				<div class="col-8 mb-1 mt-1">
					<div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Like/Dislike:</div>
				<div class="col-8 mb-1 mt-1" id="DIV_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
				</div>
			</div>
			<hr class="d-md-none my-4"/>
				
   			
   		
	   		<div class="row mb-2">
				<div class="col-12 mt-2 px-1 text-center text-footnote">
					Data provided by the SWGT Community. SWGT not responsible for inaccuracies. Also available on SWGT homepage.
				</div>
			</div>
   		</div>
   	</div>
	
	<div class="container">
  	<nav aria-label="breadcrumb">
      <ol class="breadcrumb box-shadow">
        <li class="breadcrumb-item"><a href="/">Dashboard</a></li>
        <li class="breadcrumb-item"><a>Tools</a></li>
        <li class="breadcrumb-item active" aria-current="page">Game Codes</li>
      </ol>
    </nav>
    
    <div class="card card-default box-shadow col-12 order-1 px-0">
		<div class="card-header pt-1 pb-1 pl-1 pr-1">
			<div class="row">
				<div class="col-9">
					Summoners War - Active Codes
				</div>
				<div class="col-3 text-right">
					<button type="button" class="btn btn-xs btn-primary" onClick="function_addGameCode();"><i class="fas fa-plus"></i></button>
					<button type="button" class="btn btn-xs btn-primary" onClick="function_viewModalFAQ(57);"><i class="fas fa-question"></i></button>
				</div>
			</div>
		</div>
   		<div class="card-body px-1 px-md-3">
   			
   			<div class="d-none d-md-block">
	   			<table id="activeGameCodes" class="table table-hover mb-0 table-sm font-size-08em w-100" cellspacing="0">
					<thead>
						<tr>
							<th class="text-center" style="width:20%;">Code</th>
							<th class="text-center" style="width:24%;">Date Added</th>
							<th class="text-center">Rewards</th>
							<th class="text-center" style="width:15%;">Like/Dislike</th>
						</tr>
					</thead>
					
					<tbody>
						
						<tr>
							<td>
								<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
								<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="SW2025DEC9PJ">SW2025DEC9PJ</a>
							</td>
							<td class="text-md-center">11/30/2025 03:08 PM UTC</td>
							<td><div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div></td>
							<td class="text-md-center " id="TD_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
								<button class="btn btn-xxs btn-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
							</td>
						</tr>
						
					</tbody>
				</table>
			</div>
			
				
			
			<div class="row d-md-none font-size-08em px-1">
				<div class="col-4 text-bold mb-1">Code:</div>
				<div class="col-8 mb-1">
					<button class="btn btn-xs btn-systemDefault btn-clipboard mr-1" data-clipboard-text="SW2025DEC9PJ"><i class="fas fa-clipboard" aria-hidden="true"></i></button>
					<a href="http://withhive.me/313/SW2025DEC9PJ" class="hasVisited gameCodeLink" target="_blank" data-gamecode="SW2025DEC9PJ">SW2025DEC9PJ</a>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Date Added:</div>
				<div class="col-8 mb-1 mt-1">
					11/30/2025 03:08 PM UTC
				</div>
				
				<div class="col-4 text-bold mb-1 mt-2">Rewards:</div>
				<div class="col-8 mb-1 mt-1">
					<div class="d-inline-block mr-3 mb-1" title="Energy"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/energy.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Cyrstal"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/crystal.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>100</div><div class="d-inline-block mr-3 mb-1" title="Water Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_water.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Fire Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_fire.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div><div class="d-inline-block mr-3 mb-1" title="Wind Scroll"><img src="https://do9d4mpqk497d.cloudfront.net/common/images/summoners_war_query_jp/scroll_wind.png" class="monster-image-24"/><span style="padding-right:2px;">x</span>1</div>
				</div>
				
				<div class="col-4 text-bold mb-1 mt-1">Like/Dislike:</div>
				<div class="col-8 mb-1 mt-1" id="DIV_GAME_CODE_5d9a29ae-0e9c-428f-bccd-44af0cf8740c">
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','U','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">257<i class="fas fa-thumbs-up ml-1"></i></button>
					<button class="btn btn-xxs btn-outline-primary" onclick="function_gameCodeVote('SW2025DEC9PJ','D','5d9a29ae-0e9c-428f-bccd-44af0cf8740c')">4<i class="fas fa-thumbs-down ml-1"></i></button>
				</div>
			</div>
			<hr class="d-md-none my-4"/>
				
   			
   		
	   		<div class="row mb-2">
				<div class="col-12 mt-2 px-1 text-center text-footnote">
					Data provided by the SWGT Community. SWGT not responsible for inaccuracies. Also available on SWGT homepage.
				</div>
			</div>
   		</div>
   	</div>
	`
