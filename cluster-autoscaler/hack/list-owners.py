""" Python script to list all OWNERS of various parts of Cluster Autoscaler.

Traverse all subdirectories and find OWNERS. This is useful for tagging people
on broad announcements, for instance before a new patch release.
"""

import glob
import yaml
import sys

files = glob.glob('**/OWNERS', recursive=True)
owners = set()

for fname in files:
  with open(fname) as f:
    parsed = yaml.safe_load(f)
    if 'approvers' in parsed and parsed['approvers'] is not None:
      for approver in parsed['approvers']:
        owners.add(approver)
    else:
      print("No approvers found in {}: {}".format(fname, parsed), file=sys.stderr)

for owner in sorted(owners):
  print('@', owner, sep='')
