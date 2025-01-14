# gilder-data

Generator code and JSON files containing locale and game database information acquired from public APIs and filtered 
for Gilder.

The locale and game data is found in the [data](./data) subdirectory and here in the root of the repository is an 
importer that unmarshals the JSON data and sorts it for use with Go applications.

Set your https://rawg.io API key in the environment variable `RAWG_KEY` and run the package in the [data](./data) 
directory to (re) generate the json files and gzipped versions. It will not run unless the files present there are less 
than 3 months old, as this data will not change that often. Delete the json and json.gz files in the `data` directory to 
force regeneration. Note that the games database in particular takes around 10 minutes to fully complete. 

For use with Go applications, in the root of this package these JSON files are imported and unmarshalled to package 
scoped variables `gilder_data.Nations` and `gilder_data.GamesList` and sorted alphabetically by name.