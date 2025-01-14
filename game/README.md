# games

This is a simple script that harvests multiplayer games from the [rawg games database](https://rawg.io/) and prints out a raw JSON condensed form containing the important fields for the purposes of Gilder.

You need a rawg API key to use this, create a file in this directory called `key` containing the key text, a 32 character hex string, and then this little program will pull all of the multiplayer games it can find, limited to games that have pc/linux/mac versions and no older than 2017 release.

The output should be placed into the file `games.json` in the parent directory. Optionally you can sort it using the script in `sort` in the parent directory.