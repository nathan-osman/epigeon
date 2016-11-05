FROM scratch
MAINTAINER Nathan Osman <nathan@quickmediasolutions.com>

# Add the binary
ADD dist/epigeon /usr/local/bin/epigeon

# Expose port 25 by default
EXPOSE 25

# Use epigeon as the default entrypoint
ENTRYPOINT ["/usr/local/bin/epigeon"]
