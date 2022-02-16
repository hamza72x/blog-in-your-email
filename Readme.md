# Send all of your blog posts to your email

```sh
# install the cli
$ go install -v github.com/hamza72x/blog-in-your-email@latest

# clone the repo
$ git clone https://github.com/hamza72x/blog-in-your-email && cd blog-in-your-email

# make config directory
$ mkdir -p $HOME/.config/blog-in-your-email

# copy config files to appropriate places
$ cp data.sample.csv $HOME/.config/blog-in-your-email/data.csv # make your necessary changes
$ cp config.sample.ini $HOME/.config/blog-in-your-email/config.ini # make your necessary changes

# Set cron job to run the cli (blog-in-your-email)
# Crontab example for every day at 4:00 AM
# 0 4 * * * /home/nix/go-xc/bin/blog-in-your-email > /home/nix/blog-in-your-email.log 2>&1
```
