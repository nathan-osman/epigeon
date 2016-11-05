## epigeon

Do you have an alert system or server daemon that is capable of sending alerts via email but doesn't provide a way to interact with Twitter? (I'm looking at you, Monit.) Enter epigeon - a tiny application written in Go that makes this possible.

### Installing the Application

epigeon is easy to install. Just grab the latest binary from the releases tab. The application has no dependencies and therefore is distributed as a single executable. On Linux, this is as simple as `wget`-ing the file and adding the execute bit (`chmod +x`). On Windows, the execute bit isn't even required &mdash; just run the `.exe` directly.

### Usage

The configuration for the application can be provided in one of two ways:

#### Using a File

Configuration can be provided by passing the filename of a JSON file as a command line argument. The file should look something like this:

    {
        "smtp_address": ":25",
        "twitter_consumer_key": "",
        "twitter_consumer_secret": "",
        "twitter_access_token": "",
        "twitter_access_secret": ""
    }

#### Using Environment Variables

In the absence of command line arguments, the configuration options may be specified as environment variables. The names of the environment variables are identical to the keys used in the JSON file, except that they are specified in uppercase.

### Security

You will likely want to change `smtp_address` to `127.0.0.1:25` in order to prevent external hosts from posting tweets through epigeon. Authentication will be added in a later release.
