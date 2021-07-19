# PrimeGenerator
This is my first work with concurrency and channels! This took me a lot of hard thinking but I was able to make an incredibly fast and efficient prime generator using go routines and channels!

This generator is fast because it only checks the modulo of the first half of the initial number, then THAT number is split into two variables IE:

  13 is split into 6 and 7
  
   - 0-6 is modulod to 13 and 7-13 is modulod to 13 using working that run using go routine so both will run at the same time.
    
   - Both workers send true to a waiting channel if that section of numbers has determined the number is a prime number.
    
   - If both workers send true to their waiting channels, then the number is a prime number!
