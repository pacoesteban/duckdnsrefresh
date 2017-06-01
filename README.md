duckdnsrefresh
==============

Utility to refresh dyn domain names from duckdns.org

It is pretty simple, and could be improved ... I might do that one day.

It only accepts `-v` to be a little more verbose. All the other configs need to be on a config file (YAML or JSON, it uses [viper](https://github.com/spf13/viper)). File has to be located in `$HOME/.duckdns/config.yaml` or in `.`

There's a config file example in the repo.
