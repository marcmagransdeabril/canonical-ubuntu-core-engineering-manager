# Introduction

This is a technical exercise for the mentioned role at Canonical. The exercises listed below are designed to test different skills and may require additional domain or specific knowledge, therefore don’t hesitate to send incomplete answers but in such a situation please explain your train of thought towards the solution.

# Scope & Duration

You are expected to solve both exercises. Remember your scope should only be limited to the given time frame and you may make certain assumptions unless they contradict with anything mentioned in the corresponding exercise. Please make sure to list any assumptions that you considered in your answer. 

We value your time! Please limit the scope if the solution to either question is taking more than 2 uninterrupted hours for you to complete.

# Exercise 1

Bootable Linux image via QEMU

In this exercise you are expected to create a shell script that will run in a Linux environment (will be tested on Ubuntu 20.04 LTS or 22.04 LTS). This shell script should create and run an AMD64 Linux filesystem image using QEMU that will print “hello world” after successful startup. Bonus points for creating a fully bootable filesystem image (but not mandatory). The system shouldn’t contain any user/session management or prompt for login information to access the filesystem.  

You can use any version/flavor of the Linux kernel. The script can either download and build the kernel from source on the host environment or download a publicly available pre-built kernel.

The script shouldn’t ask for any user input unless superuser privileges are necessary for some functionality, therefore any additional information that you require for the script should be available in your repository.

The script should run within the working directory and not consume any other locations on the host file system.

# Exercise 2

Shred tool in Go

Implement a Shred(path) function that will overwrite the given file (e.g. “randomfile”) 3 times with random data and delete the file afterwards. Note that the file may contain any type of data.

You are expected to give information about the possible test cases for your Shred function, including the ones that  you don’t implement,  and implementing the full test coverage is a bonus :) 

In a few lines briefly discuss the possible use cases for such a helper function as well as advantages and drawbacks of addressing them with  this approach. 

# Final Step

Create a PDF file with your answers to above questions (which consists of your repository link

URLs of your commits).

Reply to greenhouse email by attaching your answer.