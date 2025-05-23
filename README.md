# Ocelot Documentation Server

This is the source code repository for the website and documentation server, available <a href="https://ocelot-cloud.org/" target="_blank" rel="noopener noreferrer">here</a>.

Minor changes, such as fixing typos, can be made using the GitHub web interface. For larger changes, such as adding new pages, you should clone the repository and use a running Hugo server to monitor changes in real time. This will ensure correct formatting and layout of the documentation.

### Requirements

Install `hugo` and `npm`. On Debian/Ubuntu, you can do this with the following commands:

```bash
sudo apt install -y npm

latest_version=$(curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest | grep -oP '"tag_name": "\K[^"]+')
wget "https://github.com/gohugoio/hugo/releases/latest/download/hugo_extended_${latest_version#v}_Linux-64bit.tar.gz"
tar -xvzf "hugo_extended_${latest_version#v}_Linux-64bit.tar.gz"
sudo mv hugo /usr/local/bin/
```

Close the current shell and reopen it to apply the changes. Check the installation with `hugo version`.

### Setup

Clone the repository to your local machine and run `hugo server`. Go to `http://localhost:1313` in your browser to view the documentation.

### Contributing

Please read the [Community](https://ocelot-cloud.org/docs/community/) articles for more information on how to contribute to the project and interact with others.

### License

This project is licensed under the 0BSD License - see the [LICENSE](LICENSE) file for details.