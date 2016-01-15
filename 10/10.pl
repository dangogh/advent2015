#!/usr/bin/perl

my $n      = shift || "111311";
my $repeat = shift || 1;

for ( 1 .. $repeat ) {
	my $ans = '';
	while ( $n =~ /(\d)(\1)*/g ) {
		my $d = $1;
		my $p = $&;
		$ans .= length($p) . $d;
	}
	print "$ans\n";
	$n = $ans;
}
