```
$ process '{"device": "A123", "generated": "2022-01-01 15:00:00.000", "speed": 48.7, "heading": 101}'

{"device": "A123", "speed": 48.7, "heading": 101}

$ process '{"device": "A123", "generated": "2022-01-01 15:01:00.000", "position" : {"lat": -80.0101, "long": 40.0101}}'

{"device": "A123", "speed": 48.7, "heading": 101, "position" : {"lat": -80.0101, "long": 40.0101}}

$ process '{"device": "A123", "generated": "2022-01-01 15:02:00.000", "speed": 0}'

{"device": "A123", "speed": 0, "heading": 101, "position" : {"lat": -80.0101, "long": 40.0101}}

$ process '{"device": "B345", "generated": "2022-01-01 15:02:00.000", "speed": 21.55, "position" : {"lat": "-78.0101", "long": "42.0101"}}'

{"device": "B345", "speed": 21.55, "position" : {"lat": -78.0101, "long": 42.0101}}

$ process '{"device": "B345", "generated": "2022-01-01 15:01:00.000", "speed": 35.11, "heading": 45}'

{"device": "B345", "speed": 21.55, "heading": 45, "position" : {"lat": -78.0101, "long": 42.0101}}
```
