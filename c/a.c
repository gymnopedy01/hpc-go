#include <stdio.h>
int sum(int i);

main()
{
   printf("%d\n", sum(10));
}

int sum(int i)
{
   return i+sum(i-1);
}

// $gcc a.c
// $./a.out
// $Segmentation fault
