metweather(1) [""["metweather-VERSION"]]

# NAME

metweather - Display weather information from Metservice.

# SYNOPSIS

metweather <command> [flags]

# FLAGS

	*--config* _path_     
		config file (default is $XDG_CONFIG_HOME/.metweather.yaml)
	*-h, --help*
		help for metweather
	*-v, --version*
		version for metweather
	*-l, --location* _city_
		location used for the weather observation/forecast

# COMMANDS

	*observation*
		display current or past weather observations
	*forecast*
		display weather predictions for the current day or next serveral days

# AUTHORS

Maintained by Dakota Walsh <kota at nilsu.org>.
Up-to-date sources can be found at https://git.sr.ht/~kota/metweather

# COPYRIGHT

Copyright 2021 Dakota Walsh

License GPLv3+ GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it. There is NO
WARRANTY, to the extent permitted by law.
