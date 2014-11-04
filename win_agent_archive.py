#!/usr/bin/python

from __future__ import print_function

from argparse import ArgumentParser
import os
import re
import subprocess
import sys
import traceback


# The s3 container and path to add to and get from.
S3_CONTAINER = 's3://juju-qa-data/win-agents'
# The set of agents to make.
WIN_AGENT_TEMPLATES = (
    'juju-{}-win2012hvr2-amd64.tgz',
    'juju-{}-win2012hv-amd64.tgz',
    'juju-{}-win2012r2-amd64.tgz',
    'juju-{}-win2012-amd64.tgz',
    'juju-{}-win7-amd64.tgz',
    'juju-{}-win8-amd64.tgz',
    'juju-{}-win81-amd64.tgz',
)
# The versions of agent that may or will exist. The agents will
# always start with juju, the series will start with "win" and the
# arch is always amd64.
AGENT_VERSION_PATTERN = re.compile('juju-(.+)-win[^-]+-amd64.tgz')


def run(*args, **kwargs):
    command = ['s3cmd', '--no-progress']
    if 'dry_run' in kwargs:
        command.append('--dry_run')
        del kwargs['dry_run']
    if 'config' in kwargs:
        command.extend(['-c', kwargs['config']])
        del kwargs['config']
    command.extend(args)
    kwargs['stderr'] = subprocess.STDOUT
    return subprocess.check_output(command, **kwargs)


def get_source_agent_version(source_agent):
    match = AGENT_VERSION_PATTERN.match(source_agent)
    if match:
        return match.group(1)
    return None


def get_input(prompt):
    return raw_input(prompt)  # pyflakes:ignore


def listing_to_files(listing):
    agents = []
    for line in listing.splitlines():
        parts = line.split()
        agents.append(parts[-1])
    return agents


def add_agents(args):
    source_agent = os.path.basename(args.source_agent)
    version = get_source_agent_version(source_agent)
    if version is None:
        raise ValueError('%s does not look like a agent.' % source_agent)
    agent_versions = [t.format(version) for t in WIN_AGENT_TEMPLATES]
    if source_agent not in agent_versions:
        raise ValueError(
            '%s does not match an expected version.' % source_agent)
    agent_glob = '%s/juju-%s*' % (S3_CONTAINER, version)
    existing_versions = run('ls', agent_glob, config=args.config)
    if args.verbose:
        print('Checking that %s does not already exist.' % version)
    for agent_version in agent_versions:
        if agent_version in existing_versions:
            raise ValueError(
                '%s already exists. Agents cannot be overwritten.' %
                agent_version)
    # The fastest way to put the files in place is to upload the source_agent
    # then use the s3cmd cp to make remote versions.
    source_path = os.path.abspath(os.path.expanduser(args.source_agent))
    if args.verbose:
        print('Uploading %s to %s' % (source_agent, S3_CONTAINER))
    run('put', source_path, S3_CONTAINER,
        config=args.config, dry_run=args.dry_run)
    agent_versions.remove(source_agent)
    remote_source = '%s/%s' % (S3_CONTAINER, source_agent)
    for agent_version in agent_versions:
        destination = '%s/%s' % (S3_CONTAINER, agent_version)
        if args.verbose:
            print('Copying %s to %s' % (remote_source, destination))
        run('cp', remote_source, destination,
            config=args.config, dry_run=args.dry_run)


def get_agents(args):
    version = args.version
    agent_glob = '%s/juju-%s*' % (S3_CONTAINER, version)
    destination = os.path.abspath(os.path.expanduser(args.destination))
    output = run(
        'get', agent_glob, destination,
        config=args.config, dry_run=args.dry_run)
    if args.verbose:
        print(output)


def delete_agents(args):
    version = args.version
    agent_glob = '%s/juju-%s*' % (S3_CONTAINER, version)
    existing_versions = run('ls', agent_glob, config=args.config)
    if args.verbose:
        print('Checking for matching agents.')
    if version not in existing_versions:
        raise ValueError('No %s agents found.' % version)
    print(existing_versions)
    answer = get_input('Delete these versions? [y/N]')
    if answer not in ('Y', 'y', 'yes'):
        return
    agents = listing_to_files(existing_versions)
    for agent in agents:
        deleted = run('del', agent, config=args.config, dry_run=args.dry_run)
        if args.verbose:
            print(deleted)


def parse_args(args=None):
    """Return the argument parser for this program."""
    parser = ArgumentParser("Compare old and new stream data.")
    parser.add_argument(
        '-d', '--dry-run', action="store_true", default=False,
        help='Do not overwrite existing data.')
    parser.add_argument(
        '-v', '--verbose', action="store_true", default=False,
        help='Increse verbosity.')
    parser.add_argument(
        '-c', '--config', default=None, help='The s3config file.')
    subparsers = parser.add_subparsers(help='sub-command help')
    add_parser = subparsers.add_parser('add', help='Add win-agents')
    add_parser.add_argument(
        'source_agent',
        help="The win-agent to create all the agents from.")
    add_parser.set_defaults(func=add_agents)
    get_parser = subparsers.add_parser('get', help='get win-agents')
    get_parser.add_argument(
        'version', help="The version of win-agent to download")
    get_parser.add_argument(
        'destination', help="The path to download the files to.")
    get_parser.set_defaults(func=get_agents)
    get_parser = subparsers.add_parser('delete', help='get win-agents')
    get_parser.add_argument(
        'version', help="The version of win-agent to delete")
    get_parser.set_defaults(func=delete_agents)
    return parser.parse_args(args)


def main(argv):
    """Add to get win-agents."""
    args = parse_args(argv)
    try:
        args.func(args)
    except Exception as e:
        print(e)
        if args.verbose:
            traceback.print_tb(sys.exc_info()[2])
        return 2
    if args.verbose:
        print("Created mirror json.")
    return 0


if __name__ == '__main__':
    sys.exit(main(sys.argv[1:]))
