#!/usr/bin/python

import os, sys, re

path = '/Users/gdenslow/hackathon/smallset'

def parse_email( file ): 
    this_date = ''
    this_from = ''
    this_subject = ''
    f = open(path + '/' + file,'r');
    for line in f:
        date_header = '^Date:\s+(.+)'
        from_header = '^From:\s+(.+)'
        subject_header = '^Subject:\s+(.+)'
        if not this_date:
            match = re.search(date_header,line)
            this_date = match.group(1) if match else ''
        if not this_from:
            match = re.search(from_header,line)
            this_from = match.group(1) if match else ''
        if not this_subject:
            match = re.search(subject_header,line)
            this_subject = match.group(1) if match else ''
    print path + '/' + file + '|' + this_date + '|' + this_from + '|' + this_subject;

for dir_path,dirs,files in os.walk( path ):
    for file in files:
        parse_email( file )

