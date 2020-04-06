import re

def read_file(filename):
    with open(filename) as file:
        return file.read()

def get_iter_slices(given, slice_length, exclusive=True):
    given_length = len(given)
    for i in range(0,given_length if exclusive else given_length - slice_length + 1, slice_length if exclusive else 1):
        yield given[i:i+slice_length]

def parse_FASTA(contents):
    contents = contents.replace('\n','')
    content_list = re.split('>(Rosalind_\d+)', contents)
    content_list = content_list[1:]
    keys = content_list[::2]
    vals = content_list[1::2]
    ros_dict = {key:vals[i] for (i,key) in enumerate(keys)}
    return ros_dict