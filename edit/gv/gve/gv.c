/* $Id$ $Revision$ */
/* vim:set shiftwidth=4 ts=8: */

/*************************************************************************
 * Copyright (c) 2011 AT&T Intellectual Property 
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors: See CVS logs. Details at http://www.graphviz.org/
 *************************************************************************/

/* Note that, with the call to gvParseArgs(), this application assumes that
 * a known layout algorithm is going to be specified. This can be done either
 * using argv[0] or requiring the user to run this code with a -K flag specifying
 * which layout to use. In the former case, after this program has been built as
 * 'demo', you will need to rename it as one of the installed layout engines such
 * as dot, neato, sfdp, etc. 
 
 https://helpmanual.io/man8/libgvc6-config-update/
 */
#include <graphviz/gvc.h>
int
main (int argc, char **argv)
{
  GVC_t *gvc;
  graph_t *g;
  FILE *fp;
  gvc = gvContext ();
  if (argc > 1)
    fp = fopen (argv[1], "r");
  else
    fp = stdin;
  g = agread (fp, NULL);
  gvLayout (gvc, g, "dot");
  gvRender (gvc, g, "plain", stdout);
  gvFreeLayout (gvc, g);
  agclose (g);
  return (gvFreeContext (gvc));
}
