const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().write();
    try stdout.print("Hello world", .{});
}
