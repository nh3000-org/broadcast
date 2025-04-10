# Broadcast Radio Software

This system eliminates the need for a live studio and all the associated expenses related to operating a radio station

It uses technolgy deriven from a java based system to a go based system.

Some key components are:

* NATS - the storage for content
* POSTGRESQL - the database

There is a application accessed through a VPN for administration.

A small web server, also accessed through a VPN to upload a stub, and download a blank stub.

# STUB

This is just a directory structure for communicating content to the server.

Each category contains 102k MP# files with a naming convention of:

ARTIST-SONG-ALBUM.mp3

And for currents:

ARTIST-SONG-ALBUMINTRO.mp3

ARTIST-SONG-ALBUMOUTRO.mp3

Once a current moves to recurrents the INTRO and OUTRO are not played anymore and removed from the content store.

