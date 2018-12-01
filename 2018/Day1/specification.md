# Day One: Chronal Calibration

## Part One

You must calibrate the device to a certain frequency, the final frequency is calculated throug a series of frequency changes such as +3, which means an increase in frequency by 3, or -2, which signals a decrease in frequency by 2.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

+ Current frequency  0, change of +1; resulting frequency  1.
+ Current frequency  1, change of -2; resulting frequency -1.
+ Current frequency -1, change of +3; resulting frequency  2.
+ Current frequency  2, change of +1; resulting frequency  3.

In this example, the resulting frequency is 3.

Here are other example situations:

+ +1, +1, +1 results in  3
+ +1, +1, -2 results in  0
+ -1, -2, -3 results in -6

Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?