FROM scratch
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

# Add the binary
ADD dist/epigeon /usr/local/bin/

# Add the root CAs
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/

# Expose port 25 by default
EXPOSE 25

# Use epigeon as the default entrypoint
ENTRYPOINT ["/usr/local/bin/epigeon"]
