hashcheck
=========

Small little utility to check one file against multiple hashes.

Why?
====

Why use this over standard utilies like md5sum, sha1sum, etc you ask? Well keep in mind that first of all md5 is broken and you shouldn't be using it to begin with. At this point of writing the sha family isn't broken, but that might be entirely different in a couple of years. With hashcheck you can simply check against multiple hashing algorithms, it is one thing to force a hash collision with 1 algorithm, it's another thing to force a hash collision with 5 algorithms. But still why should I use this and not just use md5sum, sha1sum etc to check against multiple hashing algorithms? Well this is way shorter to write than a bash script that does all of it. Next to that hashcheck will only read your file once and not multiple times which would be the case when using a bash script to do all of it. This might not be a big deal for small files, but it is a very good thing when you're dealing with rather big files.
