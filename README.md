# aatime

aatime is a Go CLI application that displays the current time using ASCII art.

```shell
   ___     ___     ___     ___     ___                              _                _        _                    
  /   \   / __|   / __|   |_ _|   |_ _|     o O O  __ _      _ _   | |_      o O O  | |_     (_)    _ __     ___   
  | - |   \__ \  | (__     | |     | |     o      / _` |    | '_|  |  _|    o       |  _|    | |   | '  \   / -_)  
  |_|_|   |___/   \___|   |___|   |___|   TS__[O] \__,_|   _|_|_   _\__|   TS__[O]  _\__|   _|_|_  |_|_|_|  \___|  
_|"""""|_|"""""|_|"""""|_|"""""|_|"""""| {======|_|"""""|_|"""""|_|"""""| {======|_|"""""|_|"""""|_|"""""|_|"""""| 
"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-'./o--000'"`-0-0-'"`-0-0-'"`-0-0-'./o--000'"`-0-0-'"`-0-0-'"`-0-0-'"`-0-0-' 
```

## Installation

To install aatime, you need to have Go installed on your system.  
Then, you can use the following command:

```shell
go install github.com/takurooo/aatime
```

## Usage

To display the current time using ASCII art, run the following command:

```shell
aatime now [flags]
```

### help

```shell
Display current time in ASCII art.

Usage:
  aatime now [flags]

Flags:
  -f, --font string   ASCII art font. (invita or train or bulbhead) (default "invita")
  -h, --help          help for now
  -u, --utc           Display current time in UTC
```

## Examples

Display the current time using the ``invita`` font:

```shell
watch -n 1 aatime now -f invita
```

Output:

```shell
   _       __     _    _             __   _             __    ___          __      _          __     __       _        ___ 
  '  )   /   )   '  )  /   /       /   )  /   /       /   )  (   )       /   )   / /        /   )  /   )      /   /   /    
 ,--'   /   /   ,--'  /___/_  '   /   /  /___/_  '   /   /   .--'   '   /   /     /    '   /   /  (__,/   '  /___/_  /__   
/___   (__ /   /___      /    '  (__ /      /    '  (__ /   (___)   '  (__ /     /     '  (__ /      /    '     /   ____)  
                        /                  /                                    /                   /          /         
```

Display the current time using the ``train`` font:

```shell
watch -n 1 aatime now -f train
```

Output:

```shell
   ___       __      ___     _ _        _        __     _ _        _        __      ___       _        __       _        _        _        _        _       ___       __   
  |_  )     /  \    |_  )   | | |      (_)      /  \   | | |      (_)      /  \    ( _ )     (_)      /  \     / |      (_)      / |      / |      (_)     |_  )     /  \  
   / /     | () |    / /    |_  _|      _      | () |  |_  _|      _      | () |   / _ \      _      | () |    | |       _       | |      | |       _       / /     | () | 
  /___|    _\__/    /___|    _|_|_    _(_)_    _\__/    _|_|_    _(_)_    _\__/    \___/    _(_)_    _\__/    _|_|_    _(_)_    _|_|_    _|_|_    _(_)_    /___|    _\__/  
_|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| _|"""""| 
"`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-' "`-0-0-'
```

## ASCII Art Generation
The ASCII art used in this application is generated using the Text to [ASCII Art Generator](https://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20) by patorjk.com.

## License

This project is licensed under the MIT License.
