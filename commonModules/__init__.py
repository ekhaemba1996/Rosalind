import re

def read_file(filename):
    with open(filename) as file:
        return file.read()

def get_iter_slices(given, slice_length, exclusive=True):
    given_length = len(given)
    for i in range(0,given_length if exclusive else given_length - slice_length + 1, slice_length if exclusive else 1):
        yield given[i:i+slice_length]

def parse_FASTA(contents):
    pattern = re.compile('>(Rosalind_\d+)\n(([GATC]+\n)+)')
    ros_dict = {key:''.join(cap_group.split()) for key, cap_group, _ in re.findall(pattern, contents)}
    return ros_dict