from distutils.core import setup

setup(
    name='pidriver',
    version='0.0.1',
    author='Oppenheimer404',
    url='https://github.com/Oppenheimer404/PiDriver',
    packages=[
        'pidriver',
        'pidriver/database',
        'pidriver/gps',
        'pidriver/network',
        'pidriver/reporter',
        'pidriver/scanner',
    ],
    entry_points={
        'console_scripts': [
            'wifite = wifite.wifite:entry_point'
        ]
    },
    scripts=['bin/pidriver'],
)