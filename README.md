# Go RWX DLL Finder 

Tool to find DLLs with one or more RWX sections, primarily for the purpose of MockingJay process injection.

## Building

Building this project is simple. Just run:
```bash
make build
```

## Usage

To search the entire of the `C:\` drive for vulnerable DDLs, just run:
```powershell
.\rwxfinder.exe
```

To search for vulnerable DLLs under a specific path run:
```powershell
.\rwxfinder.exe --base-path <path>
```

