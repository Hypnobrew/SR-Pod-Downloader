# SR-Pod-Downloader
A tool for downloading multiple pod episodes from Sveriges Radio

1. Install the Golang bineries on your computer
2. Compile with 'go build downloader.go'
3. Run ./downloader (*nix) or downloader.exe (Windows) with the correct parameters

Use the flag --help to see required parameters

By default the downloader will run in a single thread downloading the podcasts sequently.
The downloader can run multiple worker threads by specify the flag --threads=n 