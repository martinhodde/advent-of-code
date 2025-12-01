import os

def read_file(filename):
    """
    Return a line-by-line list of contents of the given file. 
    """
    filepath = os.path.join(os.path.dirname(__file__), filename)
    file = open(filepath, 'r')
    return file.readlines()
