// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package badips

/*

Use the badIPs list
Define your security level and category

You can get the IP address list by simply using the REST API.

When you GET this URL : https://www.badips.com/get/categories
You’ll see all the different categories that are present on the service.

    Second step, determine witch score is made for you.
    Here a quote from badips that should help (personnaly I took score = 3):
    If you'd like to compile a statistic or use the data for some experiment etc. you may start with score 0.
    If you'd like to firewall your private server or website, go with scores from 2. Maybe combined with your own results, even if they do not have a score above 0 or 1.
    If you're about to protect a webshop or high traffic, money-earning e-commerce server, we recommend to use values from 3 or 4. Maybe as well combined with your own results (key / sync).
    If you're paranoid, take 5.

So now that you get your two variables, let's make your link by concatening them and grab your link.

package functions performance:

BenchmarkBadBot/first-item-access-4         	2000000000	         0.54 ns/op	1846.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkBadBot/last-item-access-4          	2000000000	         0.35 ns/op	2885.50 MB/s	       0 B/op	       0 allocs/op

*/
