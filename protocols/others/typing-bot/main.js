const robot = require("robotjs");
const { execSync } = require("child_process");

// Delay between key presses
const keyPressDelay = 50;

// Code to be typed
const code = `
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`;

// Activate VS Code
execSync("open -a 'Visual Studio Code'");

// Wait for VS Code to open
setTimeout(() => {
  // Simulate typing the code
  for (let i = 0; i < code.length; i++) {
    const char = code.charAt(i);
    robot.typeString(char);
    robot.milliSleep(keyPressDelay);
    robot.keyTap("space", "control"); // Add a small delay after each character
  }
}, 2000);
