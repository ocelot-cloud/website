# Ocelot Documentation Server

Minor changes, such as fixing typos, can be made using the GitHub web interface. For larger changes, such as adding new pages, you should clone the repository and use a running Hugo server to monitor changes in real time. This will ensure correct formatting and layout of the documentation.

### Requirements

Install hugo, go and npm. On Debian/Ubuntu, you can do this with the following commands:

```bash
sudo apt install golang-go npm

latest_version=$(curl -s https://api.github.com/repos/gohugoio/hugo/releases/latest | grep -oP '"tag_name": "\K[^"]+')
wget "https://github.com/gohugoio/hugo/releases/latest/download/hugo_extended_${latest_version#v}_Linux-64bit.tar.gz"
tar -xvzf "hugo_extended_${latest_version#v}_Linux-64bit.tar.gz"
sudo mv hugo /usr/local/bin/
hugo version
```

### Setup

Clone the repository to your local machine and run `hugo server`. Go to `http://localhost:1313` in your browser to view the documentation.

### License

This project is licensed under the 0BSD License - see the [LICENSE](LICENSE) file for details.