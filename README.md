# SeamCarving-Golang-


# Usage 
  The program takes four command line arguments
    now the program only handels png images!
    arg1: path to the input image
    arg2: path to output image 
    arg3: the new width
    arg4: the new height
    
    example:
      go run main.go img.png out.png 400 300
    
# Details
 Dynamic programming is used inorder to caculate the the energy of each seam.

# Notes
 this current version only supports image reduction. therefore if the new dimensions are equal to or larger than the   original dimensions the program will not do anything!
