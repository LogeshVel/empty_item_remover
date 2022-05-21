## empty item remover

Removes the empty files in the path you specify

This has two options 

- List the empty files in the given path and/or recursively in all directories from the given path.

- Delete all the empty files in the given path and/or recursively in all directories from the given path and then at the end lists the files status(Deleted or Not deleted) and the path.

This options can be set by the flags,
```
  -f    Provide this flag to delete the empty files
  -p string
        Path, from where the empty files need to find and proceed for deletion
  -r    Provide this flag to delete the empty Items recursively
```
If we run the program by only setting the path then by default it will lists the empty files in the given path.

-p is a mandatory flag (path flag)

If the flag(s)

- -f is set then the program will deletes and lists all the empty files in the given path.

- -r is set then the program will lists all the empty files recursively from all the directories under the given path.

- -f and -r is set then the program will deletes and lists all the empty files recursively from all the directories under the given path.

Based on the flag set to the program, we can list and/or delete the empty files from the given path and/or all other subdirectories under the given path.


  Ex:
    Consider the below folder structure in the given path.
    
    .
    |__ empty_file1
    |__ dir1
    |   |__ somefile.txt
    |__ dir2
    |   |__ empty_file2.txt
    |   |__ child_dir
    |        |__ emptyfile.py
    |         
    |__ somedir
        |__ dir
        |__ empty_file.go


To list all the empty files in the current dir which is empty_file1 in our example

```
empty-remover -p .
```

To delete and list all the empty files in the current dir which is empty_file1 in our example

```
empty-remover -p . -f
```
To list the empty files in the current dir and also in all other sub-dirs recursively which is empty_file1, empty_file2.txt, emptyfile.py, empty_file.go  in our example
```
empty-remover -p . -r
```

To delete and list all the empty files in the current dir and also in all other sub-dirs recursively which is empty_file1, empty_file2.txt, emptyfile.py, empty_file.go in our example
```
empty-remover -p . -r -f
```

### Download and Installation:
- Go to the [release page.](https://github.com/LogeshVel/empty_item_remover/releases)
- Download the ZIP file corresponding to the OS.
- Extract the ZIP file and then move the bin file to **C:/WINDOWS** in windows, **/usr/local/bin** in Linux

#### In Windows
After unzipping the zip file that you have downloaded, open the command prompt with the administrative privilege (Run as Administrator) and navigate to the path where the  **empty-remover.exe** located (extracted path) then execute the below command.
```
move empty-remover.exe "C:/WINDOWS"
```

#### In Linux
Execute the below cmd in the path where the **empty-remover** file locates
```
sudo mv empty-remover /usr/local/bin
```

_Usage based on the requirement:_
```
empty-remover -p "path/to/find"

empty-remover -p "path/to/find" -f

empty-remover -p "path/to/find" -r

empty-remover -p "path/to/find" -f -r
```