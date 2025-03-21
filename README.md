## Go & SDL2 >> 2D Platform Game Template ##
**NOTE:** That this will not work just by building, the SDL2.dll file, SDL2_image.dll file and SDL2_ttf.dll files must be in the working directory. Unfortunately the SDL2_ttf.dll file is too large to upload to GitHub so you will need to get it here https://github.com/libsdl-org/SDL_ttf/releases. Just download the package, open the archive and copy only the SDL2_ttf.dll into the working directory (where all the other files of the program are). 

Also, you need to be on the master branch of SDL2 bindings otherwise there will be build errors so once SDL2 is installed then run:<br>
```go get -v github.com/veandco/go-sdl2/sdl@master``` 

Then you should be able to build and run the program.

Created with the SDL2 bindings for Go here https://github.com/veandco/go-sdl2

Simple template for making a 2D platform game, arrow or WASD keys for movement.

Press F1 key to open the debug menu / F2 key to create a new level.

https://github.com/user-attachments/assets/4723d711-2eda-4ca1-a328-6727c6fe09bc

