# Nonogen

The goal of this project is to have a website where a user can generate a nonogram and to play it online.

The website will be written in react.js and the api used to generate the nonogram is written in Go.

You can already try the nonogam generator with your own image or simply by using the **make test** command to generate nonograms based on images from the text-pic folder

### usage :
./nonogen size brightness source (dest)
* size define the height of the nonogram. Increasing the size increase the picture details but also the time you will need to solve it
* brightness define the brightness at wich the program will consider if a square is black or white. You should tweak this value until you are satisfied of the result because it will greatly affect the complexity of the puzzle.
* source is the path to the source image. Image must be unther the jpg format.
* dest is an otptional parameter. If you precise a dest, the program will output a jpg image of the completed puzzle at the dest path


### examples :

Here are some examples of puzzles generated from the same picture, with different size and brightness parameters

Original picture

![Original Shosta](https://i.imgur.com/yndINV8.jpg)

Size 15, brightness 25000, 30000, 35000

![Original Shosta](https://i.imgur.com/KS8hl7W.jpg)
![Original Shosta](https://i.imgur.com/cy5Udgw.jpg)
![Original Shosta](https://i.imgur.com/ncmat1l.jpg)

Size 30, brightness 25000, 30000, 35000

![Original Shosta](https://i.imgur.com/uPzO8eI.jpg)
![Original Shosta](https://i.imgur.com/qMTffuA.jpg)
![Original Shosta](https://i.imgur.com/IhdnWE7.jpg)

Size 50, brightness 25000, 30000, 35000

![Original Shosta](https://i.imgur.com/dagwsJu.jpg)
![Original Shosta](https://i.imgur.com/OTl76ws.jpg)
![Original Shosta](https://i.imgur.com/DkjYfnl.jpg)

