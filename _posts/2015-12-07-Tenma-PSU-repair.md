---
layout: default
title: "Tenma 72-8695 bench PSU"
---
A long long time ago in a county the other side of the Pennines, Farnell were having a clearance of returned power supplies...

I took orders and wandered over with an empty boot and a list of names to buy as many as we could for the Manchester Hackspace (<a href="https://hacman.org.uk/">https://hacman.org.uk/</a>) and it's members. At the time I was living near Rochdale so a trip over to Leeds wasn't the two hour affair it is now :)

![Car boot loaded with PSU](/assets/2015-12-07/boot-of-car.jpg)

Wow, I didn't realise this was 2013 until I went looking for the photograph!

Anyway, when we got them back to the hackspace we found:

* Some were faulty
* Some had a fixed 5v supply in addition to the 2 variables
* Some of them were fine

We worked out which were easy fixes and then I divvied them up based on who'd asked first. People who asked early got ones that were working or that we managed to quickly repair without too many issues. There were a mixture of one of the displays not working to one side of the PSU not working.

Being the person who drove to Leeds to fetch them, bought them and then had the fun of extracting cash from people I put myself aside a fully working dual variable with the 5v supply. Boxes were labelled up and until today I didn't get to look at it again - house moves, refurbishments and life have gotten in the way of fun things recently. Today is the very first day since moving over two years ago that I've set up my bench, dragged the laptop out to write up some home automation stuff and had need of the PSU (see <a href="https://hackaday.io/project/4857-commercial-home-automation/log/20852-squeezing-a-quart-into-a-pint-pot">https://hackaday.io/project/4857-commercial-home-automation/log/20852-squeezing-a-quart-into-a-pint-pot</a> for the post in question).

So, getting back to the point......  :-D

Unboxed my PSU and it's dual PSU without 5v. Odd thinks I, thought I'd got one with all the bells. Oh crap, the right hand side of this doesn't seem to be putting out any voltage according to this display.

*check with meter*

Arse. The display isn't working but I've got voltage.

Spot the fault:

![Variable resistor with fault](/assets/2015-12-07/variable-resistor-1.jpg)

Turns out that resistor going to the terminal on the multi-turn pot is a dry joint and it's floating:

![Variable resistor with floating lead](/assets/2015-12-07/variable-resistor-2.jpg)

Check with a test lead and sure enough, that's the fault. Five minutes later, it's all soldered up and working. Huzzah!

Why do I have the one with no 5v? I think mine is at the Lancaster Hackspace (<a href="http://lamm.hackspace.org.uk/">http://lamm.hackspace.org.uk/</a>) I'll have to swap it back while nobody's looking... :)

Don't rule out buying something because it might be faulty, there's some really good bargains out there to be had!
