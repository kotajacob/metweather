# metweather [![builds.sr.ht status](https://builds.sr.ht/~kota/metweather.svg)](https://builds.sr.ht/~kota/metweather)

Display weather information from [Metservice](https://www.metservice.com/). Uses
the [metservice-go](https://git.sr.ht/~kota/metservice-go) library.

# Usage

```
metweather [command] [flags]

GLOBAL flags
      --config string     config file (default is $XDG_CONFIG_HOME/.metweather.yaml)
  -h, --help              help for metweather
  -v, --version           version for metweather
  -l, --location string   location used for the weather observation/forecast

COMMANDS
  observation: display current or past weather observations
  forecast: display weather predictions for the current day or next serveral days
```

# Config Example

```yaml
location: Dunedin
```

alternatively you can set an environment variable
`METWEATHER_LOCATION="Dunedin"`

# Resources

Discussion and patches can be found [here](https://lists.sr.ht/~kota/public-inbox).
