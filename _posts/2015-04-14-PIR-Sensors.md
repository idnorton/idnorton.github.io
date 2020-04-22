---
layout: default
title:  "Is there anybody there? Click once for yes!"
---

# Is there anybody there? Click once for yes!

So following on from my [previous post about switches](https://hackaday.io/project/4857-commercial-home-automation/log/16130-switches-are-simple-right), let’s talk about Passive Infra Red (PIR) sensors.

Previously posted on [LAMM Space](https://lamm.space/2015/04/14/is-there-anybody-there-click-once-for-yes/)

## Loxone Presence Sensor

Cost: OMG HOW MUCH?!<br/>
Summary: No way in hell I’m buying one at that price<br/>
Supplier: Loxone

![Round surface mount PIR from Loxone](/assets/2015-04-14/loxone-presence-sensor_0-300x196.png)

Wow, the price of this is staggering. It’s got a light level sensor and a PIR in it. I’m not going to buy one, but I’m including it here without much comment.
Manufacturers site – [http://shop.loxone.com/enuk/presence-sensor.html](http://shop.loxone.com/enuk/presence-sensor.html)

## Generic presence sensor

Cost: Low (£10)<br/>
Summary: Not quite what I’m after, but I can hack this.<br/>
Supplier: Low Energy Supermarket Ebay shop

![Side image of PIR](/assets/2015-04-14/generic-pir-300x212.jpg)

I purchased this with the thought that I would place them in some rooms outside of the automation system like the toilet and possibly the kitchen. I ran it in the kitchen for a while until it finally started to annoy me. There’s an adjustable timer and light level sensor which means that once the light drops below a certain level the PIR will activate the light until no presence is detected and the timer runs down. It’s 240v mains, which isn’t great for working with, so having gotten annoyed with it and removed it (Claire nearly cut herself after the lights in the kitchen went out) I took it apart for a looksee.

![PIR initial disassembly](/assets/2015-04-14/generic-pir-insides01-300x169.jpg)

So what’s the black tube? It’s a fuse, that’s a good thing. Otherwise not much of interest.

![PIR PCB before modification](/assets/2015-04-14/generic-pir-insides02-291x300.jpg)

Oh hello! That’s a 24v relay. That means that we’re looking at a 24v DC supply. *grin* The Loxone kit is 24v DC, which I don’t think I’ve mentioned so far. The two circuit boards are a mains power supply and the combined PIR and logic board with the two variable resistors on. Removing or modifying the back board *should* give us a board that we can power with 24v and will trigger a 24v line when activity and suitable light level are detected. We don’t actually care about the light level or the timer so we can turn both right down. Now we have something we can connect directly to a Loxone digital input and feed into a lighting controller block for a fraction of the cost of the official presence sensor. This I like :)

Purchased from Low Energy Supermarket Ebay Shop – [http://www.ebay.co.uk/itm/230776895665](http://www.ebay.co.uk/itm/230776895665)

## Hacking the generic PIR

Some time after writing the top half of this post, I got to the testing. Tests on the generic PIR revealed that indeed it was 24v as expected. A simple full wave rectification circuit with four 1N4007 diodes, a couple of capacitors, a signal diode and a couple of resistors.

![Components removed from PIR](/assets/2015-04-14/generic-pir-surplus-components_0-150x150.jpg)
![PIR PCB with components removed](/assets/2015-04-14/2015-04-08-19.23.58-150x150.jpg)
![PIR PCB back in housing](/assets/2015-04-14/2015-04-08-19.00.48_0-150x150.jpg)

Turns out the signal diode was mostly redundant as far as I can tell, it’s included to eliminate back EMF but that’s actually dealt with by one of the power diodes without need for it. Eh, I put one of the power diodes back in and removed the signal diode.

![PIR PCB standing vertically](/assets/2015-04-14/2015-04-08-19.38.26_1-255x300.jpg)

So testing! Lets see what the current consumption is like on this after the modification

![](/assets/2015-04-14/generic-pir-current-off-300x138.jpg)
![](/assets/2015-04-14/generic-pir-current-on-300x140.jpg)

Huzzah! 24v DC PIR with low current consumption and low cost. Our terminals are now switched 24v DC, 0v and +24v DC enabling us to use this on an input as suggested above. The input can then be mapped to the sensor input on the lighting controller and programming via the Loxone as documented.

I’m very happy with this money saving hack and will be buying a few more of these sensors in the very near future!
