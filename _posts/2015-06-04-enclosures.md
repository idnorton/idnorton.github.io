---
layout: default
title: "What shall we wrap it all up in?"
---
Any enclosure needs to be easy to access yet provide enough space to get all of our cables in and out without being too big. A tough challenge!

Once our technology solution is chosen, we have to mount it, cable it, make it all look pretty and label it. What's the best enclosure to buy for all of that? We're looking for DIN rail to mount our Loxone. Loxone are now endorsing an enclosure from Future Automation which wasn't an option when I was sourcing components:

![LNX DIN rail enclosure](/assets/2015-06-04/lnx-enclosure.jpg)

Details are at [http://www.futureautomation.co.uk/Product/Details/LXN](http://www.futureautomation.co.uk/Product/Details/LXN)

My only thoughts on reviewing this enclosure is that it has the potential to mix low voltage (Voltage Band 1) and mains (Voltage Band 2) in the same cable containment. That's dependent on what modules you install where and how you position them.

15th and 16th edition of the wiring regs makes specific mention of cables sharing the same containment (see [https://www.tlc-direct.co.uk/Book/6.6.1.htm](https://www.tlc-direct.co.uk/Book/6.6.1.htm)). We can only share containment if "every cable is insulated for the highest voltage present, or each conductor in a multicore cable is insulated for the highest voltage present". So our low voltage signalling cable for the switches has to be insulated to 230v AC in order to share containment. Loxone even recommend this in their blog post at [http://blog.loxone.com/enuk/how-to-lay-out-your-smart-home-distribution-board/](http://blog.loxone.com/enuk/how-to-lay-out-your-smart-home-distribution-board/) so if you're going to use one of these distribution boards be aware of that and make sure you follow the recommendations about keeping signalling one side and mains the other.

I chose to use an ABB Control 12755, which is not a choice I would recommend. At 36 modules and measuring in at 390 x 370 x 140mm it's also IP65. It's overkill with a capital O, and it created some of it's own interesting issues.

![ABB Control 12755 DIN rail enclosure](/assets/2015-06-04/abb-enclosure.jpg)

I opted for a large enclosure believing that it would make the task of cabling easier, but in retrospect it made it harder. The ABB range offers bus rails for this box, but good luck trying to find them to buy in the UK.

No bus rails leaves us with a couple of choices, we can add DIN mount bus rails or we can use terminal blocks, I went with terminal blocks. Again, this is not a choice I would recommend. These are the terminal components I used:

* Phoenix Contact 2715979 DIKD1.5 Through termination from MCB
* Phoenix Contact 2717016 DOK1.5 actuator for light fittings
* Phoenix Contact 2715940 Terminal block insertion bridge EB 80-DIK BU (Blue)
* Phoenix Contact 2715788 Terminal block insertion bridge EB 80-DIK WH (White)
* Phoenix Contact E/UK End clamp, DIN rail mount

![Miniserver in ABB enclosure prepared for installation](/assets/2015-06-04/miniserver-prep.jpg)

So that all looks good, what's wrong I hear you ask? Many things sadly, some of which are now too late to correct in my current installation of the miniserver. I'll talk more about the details of my cabling in a subsequent post, but for now here's a more up to date photograph:

![Miniserver in ABB enclosure installed on wall](/assets/2015-06-04/miniserver-install.jpg)

Some of the issues are:

* Separation of cabling turned out to be harder than intended, should have used smaller DIN rails and finger trunking up the sides
* I've decided I hate the terminal blocks, I don't think they make things easier
* It would be nicer to terminate the Cat5e with RJ45 plugs but that might well take more room

I attempted to save money by purchasing my 24v power supply from Farnell rather than from Loxone, which saved about £10? This turned out to be another mistake, the Farnell PSU has the terminations at the top of the enclosure, and the top of the PSU is wider than a standard DIN rail module so I had to cut the cover of the enclosure to allow for the wider power supply. It looks unseamly and the terminations are outside of the enclosure which means exposed mains when the front cover is open. I've done my best to cover this with insulated ferrules but it's not good :(

Next time I'll be blogging about expanding outputs using relay modules and the associated configuration.

