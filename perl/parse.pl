use strict;
use Data::Dumper;

my $dir = '/Users/gdenslow/hackathon/smallset';
my @files;
opendir(D,$dir) || die $!;
while( my $file = readdir D ) {
    if ( -f "$dir/$file" ) {
        push( @files, $file );
    }
}
closedir(D);

#print Dumper \@files;

for my $file ( @files ) {
    my($date_sent,$sender,$subject) = parse_email( $file );
}

sub parse_email {
    my $file = shift;
    open(F,"$dir/$file") || die $!;
    my($date_sent,$sender,$subject);
    while( my $line = <F> ) {
        if ( ! $date_sent && $line =~ /^Date:\s+(.+)/ ) {
            $date_sent = $1;
        }
        if ( ! $sender && $line =~ /^From:\s+(.+)/ ){
            $sender = $1;
        }
        if ( ! $subject && $line =~ /^Subject:\s+(.+)/ ){
            $subject = $1;
        }
    }
    close(F);
    print "$dir/$file|$date_sent|$sender|$subject\n";
}

