package main

//这一题leetCode没有go的提交，用了c++的提交
//95,94
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}

}

/*
//cpp：

 //* Definition for a binary tree node.
 //* struct TreeNode {
 //*     int val;
 //*     TreeNode *left;
 //*     TreeNode *right;
 //*     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 //* };

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        TreeNode* cur =root;
       while(1) {
            if (cur->val>p->val&&cur->val>q->val){
                cur=cur->left;
            }else if (cur->val<p->val&&cur->val<q->val){
                cur=cur->right;
            }else{
                return cur;
            }
        }
    }
};

*/
//2021.01.23 23:15 于盐城一刷完毕
//于组会和电缆网页讨论会之后
