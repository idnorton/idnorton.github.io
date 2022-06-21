---
layout: default
title: I like to keep a file of notes per day which are generally very random unordered and unstructured chunks of information, commands or code that I’m working with on that given day and occasionally downloaded files.
---
This post was originally appeared on the [Adzuna Engineering Blog](https://medium.com/@adzuna_idnorton/keeping-daily-notes-f2e83b0d41d5).

There’s three moving parts to this. The `mkdaily` script creates a dated directory structure for us and then touches the notes file:
```
#!/bin/bash
set -euo pipefail
basedir=~/daily/
date="$(date +%Y/%m/%d)"
dir=${basedir}${date}/
if [ ! -d ${dir} ]; then
    mkdir -p ${dir}
    touch ${dir}notes
fi
```

Why do we touch the notes file? Well, I like to have this in my vimrc:
```
" Turn on paste mode for notes files
autocmd BufRead notes set paste modeline ruler
```

Which means that the notes files are always in paste mode avoiding issues with indentation and formatting. Why touch the file before hand? Vim doesn’t seem to parse this setting if the file doesn’t already exist so this is a work around to make sure it does.

Then we add a shell alias to run mkdaily and then change into the newly created directory:
```
alias daily='mkdaily && cd ~/daily/$(date +%Y/%m/%d)'
```

Meaning my day usually starts with:
```
iann@iann-desktop ~
$ daily
iann@iann-desktop ~/daily/2022/02/22
$ vim notes
```

I’ve recently moved from having the directories at the top level yyyymmdd to having them split by yyyy/mm/dd which has made for a lot fewer directories at the top level and I’m happy I finally got to making that change.

Other options I’ve been recommended include:

* https://help.obsidian.md/Obsidian/Obsidian which uses Markdown

Oh and I also use paper:

* Oxford Black n’ Red A4 140 Pages Glossy Wirebound Hard-Back Ruled Notebook
* Uni-ball 751081000 UM151S SigNo Gel Rollerball Comfort Grip 0.7mm Tip 0.5mm Line Black

Paper gets used for my daily log of what I’m doing and notes from the morning standup. I use the Checklist Manifesto approach described by Adam Savage in his book Every Tool’s A Hammer: Life Is What You Make It which he picked up from his boss at Industrial Light and Magic, Brian Gernard. It’s well worth a read if you’ve not read it!
