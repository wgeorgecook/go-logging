# Go Logging
Go Logging is a wrapper around [Zap Logging](https://github.com/uber-go/zap). It abstracts the setup 
into helpful functions set to my own usual defaults. 

## Why though?
This is purely convenient for me. It might be convenient for you. I hope it will be! 
But that is not its primary concern.

# Usage
The logger must be initiated via `gologging.Start()`. You may provide a level 
if desired, the default is INFO. Be sure to call `gologging.Stop()` before your
application exits to flush the underlying logger. 

## Levels
Go Logging exposes differing log functions for each level exposed as options, DEBUG, INFO, ERROR, and FATAL. 
Calling functions lower than the provided or default level are rendered no-op. EG, if the provided level is ERROR, 
then calls to `gologging.Info()` and `gologging.Debug()` are no-op. 