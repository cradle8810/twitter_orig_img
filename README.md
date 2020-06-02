# twitter_orig_img

## 20200602 Twitter update
The source HTML file from twitter needs to using the UserAgent string which has substring 'bot'.
If you use wget(1), do like that
```
wget --user-agent='bot' <twitter tweet url>
```