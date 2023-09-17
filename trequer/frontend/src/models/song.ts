// Channel Voice Messages [nnnn = 0-15 (MIDI Channel Number 1-16)]

// 1000nnnn	0kkkkkkk
// 0vvvvvvv	Note Off event.
// This message is sent when a note is released (ended).
// (kkkkkkk) is the key (note) number. (vvvvvvv) is the velocity.

// 1001nnnn	0kkkkkkk
// 0vvvvvvv	Note On event.
// This message is sent when a note is depressed (start).
// (kkkkkkk) is the key (note) number. (vvvvvvv) is the velocity.

export type ChannelId = number;
export type PatternId = number;
export type Tick = number;

export interface MidiChannel {
  address: string;
  channel: number;
}

export interface Note {
  key: number;
  volume: number;
  length: number; //ticks
}

export interface Track {
  length: 64;
  notes: Record<Tick, Note>;
}

export interface Pattern {
  tracks: Record<ChannelId, Track>;
}

export interface Song {
  channels: Record<ChannelId, MidiChannel>;
  patterns: Record<PatternId, Pattern>;
  sequence: PatternId[];
}
