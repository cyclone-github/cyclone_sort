# Cyclone Sort
Alternative to sort (GNU coreutils) that is cross compiled for Linux, Windows & Mac. Cyclone Sort by default will sort and dedup input. See usage examples below.

Usage Examples:
- cyclone_sort -help _(displays help)_
- cyclone_sort -version _(displays version info)_
- cyclone_sortsort -i input.txt -o output.txt
- cat input.txt | cyclone_sort > output.txt
- cat input.txt | cyclone_sort _(prints to stdout)_
- _defaults to stdin if -i flag is not specified_
- _defaults to stdout if -o flag is not specified_

## Testing

### Linux:
- $ ./cyclone_sort.bin -i tmp.txt -o output.txt
- Reading file...tmp.txt
- Lines in.......16053567
- Lines out......15053568
- Dups...........999999
- Time...........5.730008666s
- Lines/Sec......2801665

### Windows:
- \> .\cyclone_sort.exe -i tmp.txt -o output.txt
- Reading file...tmp.txt
- Lines in.......16053567
- Lines out......15053568
- Dups...........999999
- Time...........7.0531078s
- Lines/Sec......2276098

### Mac:
- _untested_

### comparison between cyclone_sort, sort -u & rling
- cyclone_sort.bin -i tmp.txt -o output.txt
- 5.730s
- time sort -u tmp.txt > output.txt
- 5.434s
- time LC_ALL=C sort -u tmp.txt > output.txt
- 1.513s
- rling -s tmp.txt output.txt
- 1.039s
