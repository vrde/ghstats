**Disclaimer: this is the first time I write some Go code, if you see some :poop: you can either :see_no_evil: or open a PR and tell me how to fix my code :two_hearts:**

# gitstats
Problem: I want to measure if a software project is becoming more and more successful.
Solution: Create a tool to extract stats from a GitHub Organization, specifically how fast issues are handled, bugs addressed, pull requests merged, and what is the ratio of core contributors vs new contributors. All this data should be easy to group on a weekly/monthly/quarterly/yearly basis.

# How the Proof of Concept should look like (spoiler alert: this has not been implemented)
For the proof of concept, downloading data from GitHub and creating a CSV file should be enough. After I have the CSV file I can do all kind of queries and understand which ones are the good ones, and implement them in code.

```bash
$ gitstats -h
Usage: gitstats ORG

$ gitstats myorganization > issues.csv
```

Crazy, uh? :D
