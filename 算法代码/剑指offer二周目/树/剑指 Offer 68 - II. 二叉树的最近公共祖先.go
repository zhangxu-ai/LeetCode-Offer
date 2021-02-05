package main

func lowestCommonAncestorII(root, p,  q *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root==p||root==q{
		return root
	}
	left:=lowestCommonAncestorII(root.Left,p,q)
	right:=lowestCommonAncestorII(root.Right,p,q)
	if left==nil&&right==nil{
		return nil
	}else if left!=nil&&right!=nil{
		return root
	}else if left!=nil{
		return left
	}else{
		return right
	}
}

/*
//因为leetcode上这一题没有golang版本，所以用c++提交

 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
*/

//class Solution {
//	public:
//	TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
//		if (root==NULL){
//			return root;
//		}
//		if (root==p||root==q){
//			return root;
//		}
//		TreeNode* left =  lowestCommonAncestor(root->left,p,q);
//		TreeNode* right =  lowestCommonAncestor(root->right,p,q);
//		if (left==NULL&&right==NULL){
//			return NULL;
//		}else if (left!=NULL&&right!=NULL){
//			return root;
//		}else {
//			return left!=NULL?left:right;
//		}
//	}
//};
