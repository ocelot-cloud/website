### open todos

* add favicon to tabs, and ocelot logo to home page
* replace website by hugo
  * hugo should update automatically once a day by pulling the latest github repo stuff
* remove version creation guide in app store; reference hugo site instead
* link? "The Ocelot App Store is now live! Upload your custom apps to the App Store and install them seamlessly on any local machine via the Ocelot-Cloud."
* add content to readme.md

		var remoteHomeDir = "/home/user"
		executeOnServer("systemctl stop website")
		executeOnServer("mkdir -p %s/website", remoteHomeDir)
		rsyncCmd := fmt.Sprintf("rsync -avz website dist %s:%s/website/", sshConfigHostName, remoteHomeDir)
		tr.ExecuteInDir(backendDir, rsyncCmd)
		executeOnServer("chown -R user:user %s/website", remoteHomeDir)
		executeOnServer("chmod -R 700 %s/website", remoteHomeDir)
		executeOnServer("systemctl start website")
		executeOnServer("nmap -p 8080 localhost")