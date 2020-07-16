---
layout: default
title: "Which solution to choose?"
---
Home automation using the Loxone series of home automation products along with hacks to extend functionality and reduce cost
                
When I decided to embark on a home automation project, I spoke to a few people, read some websites and did some thinking.  You might well expect that for a project that is a substantial financial outlay and a commitment to an eco system that will be with you for a long time.

See I have a problem.  I love Open Source and free software.  I love Open Source and free hardware designs.  But I also want to be able to sell my house at some point, and in the event that I get hit by a bus I want someone else to be able to come in and understand the nest of wiring I left behind, for the sake of those left behind.

One of the electricians I spoke to told me that I couldn't build my own solution because it would cause issues with my house insurance in the event of an electrical fire.  I think that's nonsense if a sensible approach is taken and things checked over by someone suitably qualified after installation, but it's a doubt that was planted.

So for good or ill, I chose to go for a commercial off the shelf solution. A couple were recommended, I looked at them and calculated the cost.  After I'd picked myself up off the floor, I pretty much ruled the idea out on the grounds of cost.  Idly lamenting this one day at work, a colleague suggested I talk to someone in another team who was in the process of kitting out his house with some neat kit.

I spoke to Dave (a fellow IT geek), and he gave a glowing review of Loxone. It was in fact him that pointed out the issue with getting hit by a bus and someone else needing to be able to program, take over and understand whatever you leave behind. His argument was that at least with a commercial solution, worst case you can find a reseller or partner who can come in and figure it out.

So I did some reading. Looked at the specs. Looked at the cost. Looked at the cost again, then wondered how they were so much cheaper than all the other solutions. Played with the config software, downloaded the Android app. Talked to Dave some more. Then I blew £500 on a Loxone Miniserver, PSU and some toys. The spec in a nutshell:

* 8 digital inputs
* 8 digital outputs (normally open relay rated 5A)
* 4 analogue inputs (0 - 10v)
* 4 analogue outputs (0 - 10v, 20mA max)
* Onboard ethernet
* Web UI
* Free configuration software and Android application

Believe it or not, some of the commercial solutions want to charge you for the hardware and then charge you again for the damn software to configure it.  No thanks.

So the downside of this approach is that the Loxone Config software only works on Windows. Blegh. It *mostly* works under Wine, however there are some specific gotchas you need to be aware of that I'll talk about in another post.

Those are my conclusions in retrospect, they probably aren't the same as yours :)

