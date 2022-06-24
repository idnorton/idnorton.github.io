---
layout: default
title: Browser admin access to internal web interfaces using SSH and PAC

---
Throughout the career of a devops engineer, you’ll encounter many and varied web interfaces that need to be accessed in a secure way by a restricted number of people to stop Bad Things Happening.

This post was originally appeared on the [Adzuna Engineering Blog](https://medium.com/adzuna-engineering/browser-admin-access-to-internal-web-interfaces-using-ssh-and-pac-8ea201883e51).

SSH has the ability to act as a SOCKS5 proxy to enable connections to a
secure environment:
```
ssh -fN -D 8100 bastion.example.com
```

Now point your web browser proxy configuration to 127.0.0.1:8100 and you can access hosts within your remote network.

That’s useful, but then when you want to access other sites you have to change/remove your proxy config or perhaps keep two browsers, one for internal and one for public facing sites. That’s a pain.

Proxy Auto Configuration (https://en.wikipedia.org/wiki/Proxy_auto-config) or PAC solves this problem.

```
function FindProxyForURL(url, host) {
  // URLs for *.eu.local.example.com go via socks5 proxy
  if ( shExpMatch(host, "*.eu.local.example.com") ) {
    return "SOCKS localhost:8100";
  }
// All other requests go direct
  return "DIRECT";
}
```

So we can use this to direct all traffic to `*.eu.local.example.com` to the SSH proxy and all other requests should be sent direct! Assuming that the hostnames resolve to private addresses which is a whole different discussion.

If you have more than one environment you need to connect to then things get more complicated, but that’s okay because we can have multiple SSH sessions:
```
ssh -fN -D 8100 bastion-au.example.com
ssh -fN -D 8101 bastion-eu.example.com
ssh -fN -D 8102 bastion-us.example.com
```

And now we can make that PAC more complicated:
```
function FindProxyForURL(url, host) {
  // URLs for *.au.local.example.com go via socks5 proxy
  if ( shExpMatch(host, "*.au.local.example.com") ) {
    return "SOCKS localhost:8100";
  }
// URLs for *.eu.local.example.com go via socks5 proxy
  if ( shExpMatch(host, "*.eu.local.example.com") ) {
    return "SOCKS localhost:8101";
  }
// URLs for *.us.local.example.com go via socks5 proxy
  if ( shExpMatch(host, "*.us.local.example.com") ) {
    return "SOCKS localhost:8102";
  }
// All other requests go direct
  return "DIRECT";
}
```

This is not a technique I use now as I’ve switched from Firefox to Chrome and under Linux this is less convenient than it really should be. Alternative approaches include making use of proxy configuration plugin, or if you’re using Amazon Web Services then you might consider setting up Cognito services to give you authenticated access to your internal interfaces which is the way to go for less technical end users or organisations with a Single Sign On setup.

None of this is new, but it’s nice to have a note of all the details in one place.

Thanks to Adam Taylor
