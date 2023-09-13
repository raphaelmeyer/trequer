// Channel Voice Messages [nnnn = 0-15 (MIDI Channel Number 1-16)]

// 1000nnnn	0kkkkkkk
// 0vvvvvvv	Note Off event.
// This message is sent when a note is released (ended).
// (kkkkkkk) is the key (note) number. (vvvvvvv) is the velocity.

// 1001nnnn	0kkkkkkk
// 0vvvvvvv	Note On event.
// This message is sent when a note is depressed (start).
// (kkkkkkk) is the key (note) number. (vvvvvvv) is the velocity.

// song
// track
// sequence
// row
// pattern
// channel
// note
// order

export interface Channel {
  address: string;
  channel: number;
}

export interface Note {
  tick: number;
  key: number;
  volume: number;
  length: number; //ticks
}

export interface Track {
  notes: Note[];
}

export interface Pattern {
  tracks: Track[];
}

export interface Song {
  channels: Channel[];
  patterns: Pattern[];
  sequence: number[];
}
