const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const stdin = std.io.getStdIn().reader();

    // Print hello message
    try stdout.print("Hello world on ZIG\n", .{});
    try stdout.print("Please write something: ", .{});

    // Create a buffer to read input
    var buffer: [100]u8 = undefined;

    // Read a line from stdin
    if (try stdin.readUntilDelimiterOrEof(buffer[0..], '\n')) |input| {
        try stdout.print("You wrote: {s}\n", .{input});
    } else {
        try stdout.print("No input received\n", .{});
    }
}
