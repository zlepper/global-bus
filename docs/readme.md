# Global Bus

A library for building a Global Service Bus that is available in multiple
distinct runtimes at the same time, all over the same bus, with translation
between runtimes handled automatically. 

In addition, Global Bus aims for be pretty damn fast.

Global Bus is inspired by [NServiceBus](https://github.com/Particular/NServiceBus)
from Particular. 

Overall the goals of Global Bus is the following:
* Be developer friendly
* Be runtime agnostic
* Be fast
* Be free

## How?
The messages are defined in the Google Protobuf format, thus allowing them 
to be both created and read by different runtimes. 

# Contributing

## Project structure

[docs](docs/): Contains all documentation for the project.  
[lib](lib/): Contains the shared library code used by all Global Bus runtimes.  
[samples](samples/): Contains sample in each language in the main repository  
[scripts](scripts/): Contains scripts for doing different things in the repository  
[include](include/): Various external files used by this project, but not
changed by it. e.g., the main google protobuf base files.   

