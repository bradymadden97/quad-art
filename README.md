# Go-Quads

Pixel art using quad trees in Go. Inspired by https://github.com/fogleman/Quads. 
<br><br>
## Examples
![Imgur](http://i.imgur.com/ykwp2Aj.jpg)<br><br>
![Imgur](http://i.imgur.com/glXj1BJ.jpg)<br><br>
![Imgur](http://i.imgur.com/zZ5s31K.jpg)<br><br>
![Imgur](https://i.imgur.com/3pizpO7.jpg)


## Usage
` -f <filename> ` : Input image filename

` -i <iterations> ` : Number of iterations to run quads - default 200

` -b ` : Add borders to subimages

` -bc <R,G,B> ` : Border/ background color between subimages - default 0,0,0

` -c ` : Modify quads to circles

` -ds ` : Don't save intermediate images

` -g ` : Flag to create gif of quad images

` -gf <frames> ` : Number of images per second in gif - default 20

` -gp <pause> ` : Number of seconds to pause at end of gif - default 2
