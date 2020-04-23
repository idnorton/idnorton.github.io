---
layout: default
title:  "Posts"
---
<ul>
  {% for post in site.posts %}
    <div>
      <h2><a href="{{ post.url }}">{{ post.title }}</a></h2>
      <i>{{ post.date | date: "%Y-%m-%d" }}</i>
      {{ post.excerpt }}
      <hr />
    </div>
  {% endfor %}
</ul>
