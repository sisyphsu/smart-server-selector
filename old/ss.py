#!/usr/bin/env python2

from __future__ import print_function

import itertools
import shlex
import sys
from collections import namedtuple
from subprocess import Popen, PIPE

Response = namedtuple('Response', 'returncode value')


def flatten(data):
    return list(itertools.chain.from_iterable(data))


class Whiptail(object):
    def __init__(self, title='', backtitle='', height=10, width=50, auto_exit=True):
        self.title = title
        self.backtitle = backtitle
        self.height = height
        self.width = width
        self.auto_exit = auto_exit

    def run(self, control, msg, extra=(), exit_on=(1, 255)):
        whiptail = 'dialog'
        cmd = [
            whiptail, '--title', self.title, '--backtitle', self.backtitle,
            '--' + control, msg, str(self.height), str(self.width)
        ]
        cmd += list(extra)
        p = Popen(cmd, stderr=PIPE)
        out, err = p.communicate()
        if self.auto_exit and p.returncode in exit_on:
            print('User cancelled operation.')
            sys.exit(p.returncode)
        return Response(p.returncode, err)

    def prompt(self, msg, default='', password=False):
        control = 'passwordbox' if password else 'inputbox'
        return self.run(control, msg, [default]).value

    def confirm(self, msg, default='yes'):
        defaultno = '--defaultno' if default == 'no' else ''
        return self.run('yesno', msg, [defaultno], [255]).returncode == 0

    def alert(self, msg):
        self.run('msgbox', msg)

    def view_file(self, path):
        self.run('textbox', path, ['--scrolltext'])

    def calc_height(self, msg):
        height_offset = 8 if msg else 7
        return [str(self.height - height_offset)]

    def menu(self, msg='', items=(), prefix=' - '):
        if isinstance(items[0], str):
            items = [(i, '') for i in items]
        else:
            items = [(k, prefix + v) for k, v in items]
        extra = self.calc_height(msg) + flatten(items)
        return self.run('menu', msg, extra).value

    def showlist(self, control, msg, items, prefix):
        if isinstance(items[0], str):
            items = [(i, '', 'OFF') for i in items]
        else:
            items = [(k, prefix + v, s) for k, v, s in items]
        extra = self.calc_height(msg) + flatten(items)
        return shlex.split(self.run(control, msg, extra).value)

    def radiolist(self, msg='', items=(), prefix=' - '):
        return self.showlist('radiolist', msg, items, prefix)

    def checklist(self, msg='', items=(), prefix=' - '):
        return self.showlist('checklist', msg, items, prefix)


# from termcolor import colored


# load config fail
def load_servers(filename):
    result = []
    with open(filename, 'r') as _servers:
        lines = _servers.readlines()
        for line in lines:
            line = line.strip()
            if len(line) == 0 or line[0] == '#' or line[0] == '/':
                continue

            if line.count(',') < 2:
                print('invalid: ' + line)
                continue

            env = line[0:line.index(',')]
            line = line[line.index(',') + 1:]
            addr = line[0:line.index(',')]
            line = line[line.index(',') + 1:]
            desc = line

            result.append({
                'env': env,
                'addr': addr,
                'desc': desc
            })
    return result


# show server selector, based on whiptail or dialog
def select_server(servers):
    Whiptail().radiolist("choice", [str('test,1235'), str('pre,12345')], 'pre')
    return


if __name__ == '__main__':
    # os.system("ssh -p8765 sull@13.127.159.17")
    servers = load_servers("servers")
    select_server(servers)
